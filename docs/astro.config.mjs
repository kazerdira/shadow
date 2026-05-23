// @ts-check

import starlight from "@astrojs/starlight";
import starlightClientMermaid from "@pasqal-io/starlight-client-mermaid";
import { defineConfig } from "astro/config";
import starlightLlmsTxt from "starlight-llms-txt";
import starlightLinksValidator from "starlight-links-validator";

// https://astro.build/config
export default defineConfig({
	site: "https://shadow-docs.kazerdira.me",
	integrations: [
		starlight({
			title: "shadow Robot",
			plugins: [
				starlightClientMermaid(),
				starlightLlmsTxt(),
				starlightLinksValidator(),
			],
			social: [
				{
					icon: "github",
					label: "GitHub",
					href: "https://github.com/kazerdira/shadow",
				},
			],
			sidebar: [
				{
					label: "Getting Started",
					items: [
						{ label: "Introduction", slug: "getting-started/introduction" },
						{ label: "Quick Start", slug: "getting-started/quick-start" },
					],
				},
				{
					label: "Commands",
					collapsed: false,
					autogenerate: { directory: "commands" },
				},
				{
					label: "Self-Hosting",
					collapsed: true,
					autogenerate: { directory: "self-hosting" },
				},
				{
					label: "Architecture",
					collapsed: true,
					autogenerate: { directory: "architecture" },
				},
				{
					label: "API Reference",
					collapsed: true,
					autogenerate: { directory: "api-reference" },
				},
				{
					label: "Contributing",
					collapsed: true,
					autogenerate: { directory: "contributing" },
				},
			],
		}),
	],
});
