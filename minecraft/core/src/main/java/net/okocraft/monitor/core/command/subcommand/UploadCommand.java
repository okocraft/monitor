package net.okocraft.monitor.core.command.subcommand;

import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.error.EncodeError;
import dev.siroshun.codec4j.api.io.ElementAppender;
import dev.siroshun.codec4j.api.io.Out;
import dev.siroshun.codec4j.io.gson.GsonIO;
import dev.siroshun.jfun.result.Result;
import net.okocraft.monitor.core.cloud.data.ObjectMeta;
import net.okocraft.monitor.core.cloud.data.SignedData;
import net.okocraft.monitor.core.cloud.sign.HmacSigner;
import net.okocraft.monitor.core.cloud.storage.CloudStorage;
import net.okocraft.monitor.core.cloud.storage.UploadError;
import net.okocraft.monitor.core.command.Command;
import net.okocraft.monitor.core.command.CommandSender;
import net.okocraft.monitor.core.logger.MonitorLogger;
import net.okocraft.monitor.core.models.data.PlayerConnectLogData;
import net.okocraft.monitor.core.storage.PlayerLogStorage;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.UnknownNullability;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.time.Instant;
import java.time.temporal.ChronoUnit;
import java.util.ArrayList;
import java.util.Base64;
import java.util.List;
import java.util.UUID;

public class UploadCommand extends AbstractLookupCommand implements Command {

    private final CloudStorage cloudStorage;
    private final HmacSigner signer;

    public UploadCommand(PlayerLogStorage storage, CloudStorage cloudStorage, HmacSigner signer) {
        super(storage);
        this.cloudStorage = cloudStorage;
        this.signer = signer;
    }

    @Override
    public void execute(CommandSender sender, String[] args) {
        List<PlayerConnectLogData> logs = new ArrayList<>();

        try {
            this.lookupConnectLog(logs::add);
        } catch (Exception e) {
            sender.sendPlainMessage("Failed to lookup connect log: " + e.getMessage());
            return;
        }

        if (logs.isEmpty()) {
            sender.sendPlainMessage("No connect log found.");
            return;
        }

        UUID id = UUID.randomUUID();
        ObjectMeta meta = new ObjectMeta(id, ObjectMeta.ObjectType.PLAYER_CONNECT_LOG, ObjectMeta.CURRENT_VERSION, Instant.now().plus(7, ChronoUnit.DAYS));
        Result<SignedData<ObjectMeta>, EncodeError> singedMetaResult = this.signer.sign(meta, ObjectMeta.ENCODER);
        if (singedMetaResult.isFailure()) {
            sender.sendPlainMessage("Failed to create meta.");
            MonitorLogger.logger().error("Failed to create meta: {}", singedMetaResult.unwrapError());
            return;
        }

        SignedData<ObjectMeta> singedMeta = singedMetaResult.unwrap();
        byte[] metaData;
        try (ByteArrayOutputStream out = new ByteArrayOutputStream()) {
            Result<Void, EncodeError> result = GsonIO.DEFAULT.encodeTo(out, SignedData.ENCODER_WITHOUT_META, singedMeta);
            if (result.isFailure()) {
                sender.sendPlainMessage("Failed to encode meta.");
                MonitorLogger.logger().error("Failed to encode meta: {}", result.unwrapError());
                return;
            }
            metaData = out.toByteArray();
        } catch (IOException e) {
            throw new RuntimeException(e); // should not reach here.
        }

        Result<Void, UploadError> uploadResult = this.cloudStorage.upload("minecraft/logs/" + id, new Encoder<>() {
            @Override
            public @NotNull <O> Result<O, EncodeError> encode(@NotNull Out<O> out, @UnknownNullability List<PlayerConnectLogData> playerConnectLogData) {
                Result<ElementAppender<O>, EncodeError> appender = out.createList();
                if (appender.isFailure()) {
                    return appender.asFailure();
                }
                for (PlayerConnectLogData log : playerConnectLogData) {
                    Result<O, EncodeError> appendResult = appender.unwrap().append(elementOut -> PlayerConnectLogData.ENCODER.encode(elementOut, log));
                    if (appendResult.isFailure()) {
                        return appendResult.asFailure();
                    }
                }
                return appender.unwrap().finish();
            }
        }, logs);

        if (uploadResult.isFailure()) {
            sender.sendPlainMessage("Failed to upload connect log.");
            MonitorLogger.logger().error("Failed to upload connect log: {}", uploadResult.unwrapError());
            return;
        }

        String metaQuery = Base64.getUrlEncoder().encodeToString(metaData);

        sender.sendPlainMessage("Upload finished (" + logs.size() + " logs)");
        sender.sendPlainMessage("Viewer url: https://example.com/logs/view/" + id + "?meta=" + URLEncoder.encode(metaQuery, StandardCharsets.UTF_8));
    }
}
