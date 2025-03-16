import { Text } from "../../components/ui/Text";
import { Title } from "../../components/ui/Title";
import { useAuth } from "../../hooks/useAuth.ts";

export const Component = () => {
    const auth = useAuth();
    const me = auth.me.current;

    if (!me || !me.nickname) {
        throw new Error("unauthorized");
    }

    return (
        <div className="m-5">
            <Title type="large" className="mb-3">
                My page
            </Title>
            <Text>
                You are logged in as <b>{me.nickname}</b>.
            </Text>
            <Title type="base" className="my-2">
                Account
            </Title>
            <ul>
                <li>
                    <Text>ID: {me.uuid}</Text>
                </li>
                <li>
                    <Text>Nickname: {me.nickname}</Text>
                </li>
                <li>
                    <Text>Role: {me.roleName}</Text>
                </li>
            </ul>
        </div>
    );
};
