import {createFileRoute, redirect} from '@tanstack/react-router'
import {Header} from "../../../components/ui/Header";

export const Route = createFileRoute('/settings/roles/')({
	beforeLoad: async ({ context }) => {
		const isAuthenticated = await context.auth.isAuthenticated();
		if (!isAuthenticated || !context.auth.pagePermission.current?.settings.roles) {
			throw redirect({
				to: "/",
			});
		}
	},
	component: RouteComponent,
})


function RouteComponent() {
	return (
		<>
			<Header />
			<div>Hello "/settings/roles/"!</div>
		</>
	);
}
