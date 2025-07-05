package net.okocraft.monitor.core.command.subcommand;

import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.error.EncodeError;
import dev.siroshun.codec4j.io.base64.Base64IO;
import dev.siroshun.codec4j.io.gson.GsonIO;
import dev.siroshun.jfun.result.Result;
import net.okocraft.monitor.core.cloud.data.ObjectMeta;
import net.okocraft.monitor.core.cloud.data.SignedData;
import net.okocraft.monitor.core.cloud.sign.HmacSigner;
import net.okocraft.monitor.core.cloud.storage.CloudStorage;
import net.okocraft.monitor.core.cloud.storage.UploadError;
import net.okocraft.monitor.core.command.CommandSender;
import net.okocraft.monitor.core.command.subcommand.lookup.LogLookup;
import net.okocraft.monitor.core.command.subcommand.lookup.ParamParseError;
import net.okocraft.monitor.core.logger.MonitorLogger;
import net.okocraft.monitor.core.models.data.PlayerChatLogData;
import net.okocraft.monitor.core.models.data.PlayerConnectLogData;
import net.okocraft.monitor.core.models.data.UploadedObject;
import net.okocraft.monitor.core.models.lookup.CommonLookupParams;
import net.okocraft.monitor.core.storage.PlayerLogStorage;
import net.okocraft.monitor.core.storage.UploadedObjectStorage;

import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.sql.SQLException;
import java.time.Instant;
import java.time.LocalDateTime;
import java.time.temporal.ChronoUnit;
import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.UUID;

public class UploadCommand extends AbstractLookupCommand {

    private final UploadedObjectStorage uploadedObjectStorage;
    private final CloudStorage cloudStorage;
    private final HmacSigner signer;

    public UploadCommand(PlayerLogStorage storage, UploadedObjectStorage uploadedObjectStorage, CloudStorage cloudStorage, HmacSigner signer) {
        super(storage, 10000);
        this.uploadedObjectStorage = uploadedObjectStorage;
        this.cloudStorage = cloudStorage;
        this.signer = signer;
    }

    @Override
    public void execute(CommandSender sender, String[] args) {
        if (args.length < 2) {
            sender.sendPlainMessage("Usage: /monitor upload <type> {params}");
            return;
        }

        switch (args[1].toLowerCase()) {
            case "connect" ->
                this.lookupAndUpload(sender, args, this.connectLogLookup, ObjectMeta.ObjectType.PLAYER_CONNECT_LOG, PlayerConnectLogData.ENCODER.toCollectionEncoder());
            case "chat" ->
                this.lookupAndUpload(sender, args, this.chatLogLookup, ObjectMeta.ObjectType.PLAYER_CHAT_LOG, PlayerChatLogData.ENCODER.toCollectionEncoder());
            default -> sender.sendPlainMessage("Unknown type: " + args[1]);
        }
    }

    private <P extends CommonLookupParams, T> void lookupAndUpload(CommandSender sender, String[] args, LogLookup<P, T> lookup, ObjectMeta.ObjectType objectType, Encoder<List<T>> encoder) {
        Result<P, ParamParseError> paramResult = this.parseAsLookupParams(args, lookup);
        if (paramResult.isFailure()) {
            sender.sendPlainMessage(paramResult.unwrapError().message());
            return;
        }

        P params = paramResult.unwrap();

        List<T> logs = new ArrayList<>();
        try {
            lookup.lookupByParam().lookup(params, logs::add);
        } catch (SQLException e) {
            sender.sendPlainMessage("Failed to lookup " + lookup.type() + " log: " + e.getMessage());
            MonitorLogger.logger().error("Failed to lookup {} log", lookup.type(), e);
        }

        if (logs.isEmpty()) {
            sender.sendPlainMessage("No " + lookup.type()  + " log found.");
            return;
        }

        this.upload(sender, objectType, encoder, logs);
    }

    private <T extends Collection<?>> void upload(CommandSender sender, ObjectMeta.ObjectType type, Encoder<T> encoder, T logs) {
        UUID id = UUID.randomUUID();
        ObjectMeta meta = new ObjectMeta(id, type, ObjectMeta.CURRENT_VERSION, Instant.now().plus(7, ChronoUnit.DAYS));
        Result<SignedData<ObjectMeta>, EncodeError> singedMetaResult = this.signer.sign(meta, ObjectMeta.ENCODER);
        if (singedMetaResult.isFailure()) {
            sender.sendPlainMessage("Failed to create meta.");
            MonitorLogger.logger().error("Failed to create meta: {}", singedMetaResult.unwrapError());
            return;
        }

        Result<String, EncodeError> metaQueryResult =
            Base64IO.createUrlBase64(GsonIO.DEFAULT)
                .encodeToBytes(SignedData.ENCODER_WITHOUT_META, singedMetaResult.unwrap())
                .map(data -> new String(data, StandardCharsets.UTF_8))
                .map(data -> URLEncoder.encode(data, StandardCharsets.UTF_8));
        if (metaQueryResult.isFailure()) {
            sender.sendPlainMessage("Failed to create meta query.");
            MonitorLogger.logger().error("Failed to create meta query: {}", metaQueryResult.unwrapError());
            return;
        }

        try {
            this.uploadedObjectStorage.recordUploadedObject(new UploadedObject(
                id, meta.type().ordinal(), meta.version(), sender.uuid(), sender.name(), LocalDateTime.now(), meta.expiresAt()
            ));
        } catch (SQLException e) {
            sender.sendPlainMessage("Failed to record uploaded object: " + e.getMessage());
            MonitorLogger.logger().error("Failed to record uploaded object", e);
            return;
        }

        Result<Void, UploadError> uploadResult = this.cloudStorage.upload("minecraft/logs/" + id, encoder, logs);
        if (uploadResult.isFailure()) {
            sender.sendPlainMessage("Failed to upload connect log.");
            MonitorLogger.logger().error("Failed to upload connect log: {}", uploadResult.unwrapError());
            return;
        }

        sender.sendPlainMessage("Upload finished (" + logs.size() + " logs)");
        sender.sendPlainMessage("Viewer url: https://example.com/logs/view/" + id + "?meta=" + metaQueryResult.unwrap());
    }
}
