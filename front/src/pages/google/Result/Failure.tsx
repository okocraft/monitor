import { ExclamationTriangleIcon } from "@heroicons/react/24/solid";
import { Link } from "@tanstack/react-router";
import { Text } from "../../../components/ui/Text";
import { Title } from "../../../components/ui/Title";

export type Props = {
    title: string;
    description: string;
};

export const Failure = (props: Props) => {
    return (
        <div className="m-5">
            <Title color="red" className="flex">
                <ExclamationTriangleIcon className="w-10 h-10 mr-1" />
                <span className="my-auto">{props.title}</span>
            </Title>
            <div className="my-3 ml-3">
                <Text>{props.description}</Text>
            </div>
            <Link to="/" className="ml-3">
                <Text color="primary" hoverColor="primaryDark">
                    Back to top page
                </Text>
            </Link>
        </div>
    );
};
