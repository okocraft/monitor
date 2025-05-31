import type { Meta } from "./meta";
import { errorResponse, renderGZip } from "./response";
import { type SignedJsonData, createCryptoKey, verifyMeta } from "./sign";

export default {
	async fetch(request, env): Promise<Response> {
		const debug = env.DEBUG === "true";

		if (request.method !== "GET") {
			return errorResponse(405, `Method not allowed: ${request.method}`, debug);
		}

		const secretKey = env.SECRET_KEY;
		const bucket = env.MONITOR_OBJECT_BUCKET;

		if (!secretKey || !bucket) {
			return errorResponse(
				500,
				"some environment variables are not set",
				debug,
			);
		}

		const key = await createCryptoKey(secretKey);

		const url = new URL(request.url);
		const metaParam = url.searchParams.get("meta");
		if (!metaParam) {
			return errorResponse(400, "Missing meta parameter", debug);
		}

		let meta: Meta;
		try {
			const metaObj = JSON.parse(atob(metaParam)) as SignedJsonData;
			if (!metaObj.data || !metaObj.signature) {
				return errorResponse(400, "Invalid meta format", debug);
			}
			if (!(await verifyMeta(metaObj, key))) {
				return errorResponse(400, "Invalid meta signature", debug);
			}
			meta = JSON.parse(metaObj.data) as Meta;
		} catch (e) {
			return errorResponse(400, "Invalid meta JSON", debug);
		}

		if (meta.expires_at < Date.now()) {
			return errorResponse(410, "Meta is expired", debug);
		}

		const object = await bucket.get(`minecraft/logs/${meta.id}`);
		if (!object || !object.body) {
			return errorResponse(404, `Object not found for id ${meta.id}`, debug);
		}

		return renderGZip(object.body);
	},
} satisfies ExportedHandler<Env>;
