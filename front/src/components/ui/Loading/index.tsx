import { Spin } from "../Spin";
import { Text } from "../Text";

export const Loading = () => {
    return (
        <div className="flex m-auto min-h-full">
            <Spin />
            <Text type="large" className="ml-2 my-auto">
                Loading...
            </Text>
        </div>
    );
};
