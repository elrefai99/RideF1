import { Container } from "@cloudflare/containers";

interface Env {
	GIN_CONTAINER: DurableObjectNamespace<GinContainer>;
}

export class GinContainer extends Container {
	defaultPort = 8080;
}

export default {
	async fetch(request: Request, env: Env): Promise<Response> {
		const container = env.GIN_CONTAINER.getByName("gin-api");
		return container.fetch(request);
	},
};
