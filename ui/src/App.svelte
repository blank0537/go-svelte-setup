<script>
	// @ts-nocheck

	import { Link } from "svelte-routing";
	import RouterWrap from "./Routes.svelte";
</script>

<main>
	<RouterWrap>
		<nav>
			<Link to="/">Home</Link>
			<Link to="/test">Test 1</Link>
			<Link to="/test/1">Test 2</Link>
			<Link to="/invalid">Error</Link>
		</nav>
	</RouterWrap>
</main>

<style>
</style>
