export type SignedJsonData = {
	data: string;
	signature: string;
};

const keyCache: { key: CryptoKey | null; secretKey: string } = {
	key: null,
	secretKey: "",
};

export async function createCryptoKey(secretKey: string) {
	if (keyCache.key !== null && keyCache.secretKey === secretKey) {
		return keyCache.key;
	}

	const key = await importHMacSecretKey(secretKey);
	keyCache.key = key;
	keyCache.secretKey = secretKey;
	return key;
}

async function importHMacSecretKey(key: string) {
	const keyData = encodeText(key);
	return await crypto.subtle.importKey(
		"raw",
		keyData,
		{ name: "HMAC", hash: "SHA-256" },
		false,
		["sign"],
	);
}

export async function verifyMeta(
	meta: SignedJsonData,
	key: CryptoKey,
): Promise<boolean> {
	const signatureBuffer = await crypto.subtle.sign(
		"HMAC",
		key,
		encodeText(meta.data),
	);

	const expectedHex = Array.from(new Uint8Array(signatureBuffer))
		.map((b) => b.toString(16).padStart(2, "0"))
		.join("");

	if (expectedHex.length !== meta.signature.length) return false;

	let result = 0;
	for (let i = 0; i < expectedHex.length; i++) {
		result |= expectedHex.charCodeAt(i) ^ meta.signature.charCodeAt(i);
	}
	return result === 0;
}

const encoder = new TextEncoder();

export function encodeText(text: string) {
	return encoder.encode(text);
}
