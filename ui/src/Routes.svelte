<script>
	// @ts-nocheck

	import Home from "./pages/Home.svelte";
	import Test from "./pages/Test.svelte";
	import Counter from "./lib/Counter.svelte";
	import Error from "./pages/Error.svelte";
	import { Router, Route } from "svelte-routing";

	let url = "";
</script>

<Router {url}>
	<slot />
	<Route path="/" component={Home} />
	<Route path="/test" component={Test} />
	<Route path="/test/:id" component={Counter} />
	<Route path="*" component={Error} />
</Router>
