export type TruncatedTextProps = {
    text: string;
    maxLength: number;
};

export const TruncatedText = (props: TruncatedTextProps) => {
    return truncateTo16Alphabets(props.text, props.maxLength);
};

const truncateTo16Alphabets = (text: string, maxLength: number) => {
    let bytes = 0;
    let truncated = "";

    for (let i = 0; i < text.length; i++) {
        const char = text.codePointAt(i);
        if (!char) {
            break;
        }

        bytes += 0x20 <= char && char <= 0x7e ? 1 : 2;
        if (maxLength < bytes) {
            truncated += "...";
            break;
        }
        truncated += text.charAt(i);
    }

    return truncated;
};
