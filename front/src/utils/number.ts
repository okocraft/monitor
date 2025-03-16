const MAX_INT64 = (1n << 63n) - 1n;

export function isValidInt64ForHex(value: string): boolean {
    try {
        const num = BigInt(`0x${value}`);
        return num >= 0n && num < MAX_INT64;
    } catch {
        return false;
    }
}
