<script lang="ts">
	const features = [
		{
			title: "Database Connections",
			description: "Securely manage and store your database connections for quick access.",
			icon: "🔌"
		},
		{
			title: "Visual Query Builder",
			description: "Build complex SQL queries through an intuitive visual interface.",
			icon: "🔧"
		},
		{
			title: "Data Visualization",
			description: "Visualize your query results with interactive charts and graphs.",
			icon: "📊"
		}
	];
</script>

<svelte:head>
	<title>QB | Query Builder</title>
	<meta name="description" content="A visual database query builder" />
</svelte:head>

<div class="hero">
	<h1>Query Builder</h1>
	<p class="subtitle">Build, test, and visualize database queries with ease</p>
</div>

<div class="features">
	{#each features as feature}
		<div class="feature-card">
			<span class="icon">{feature.icon}</span>
			<h2>{feature.title}</h2>
			<p>{feature.description}</p>
		</div>
	{/each}
</div>

<div class="cta">
	<a href="/connections" class="button">Get Started</a>
</div>

<style>
	.hero {
		text-align: center;
		padding: 4rem 2rem;
	}

	h1 {
		font-size: 3rem;
		font-weight: bold;
		margin-bottom: 1rem;
		color: var(--color-theme-1);
	}

	.subtitle {
		font-size: 1.5rem;
		color: var(--color-text-2);
		max-width: 600px;
		margin: 0 auto;
	}

	.features {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
		gap: 2rem;
		padding: 2rem;
		max-width: 1200px;
		margin: 0 auto;
	}

	.feature-card {
		background: var(--color-bg-1);
		padding: 2rem;
		border-radius: 8px;
		text-align: center;
		transition: transform 0.2s;
	}

	.feature-card:hover {
		transform: translateY(-5px);
	}

	.icon {
		font-size: 3rem;
		margin-bottom: 1rem;
		display: block;
	}

	.feature-card h2 {
		font-size: 1.5rem;
		margin-bottom: 1rem;
		color: var(--color-theme-1);
	}

	.feature-card p {
		color: var(--color-text-2);
		line-height: 1.5;
	}

	.cta {
		text-align: center;
		padding: 4rem 2rem;
	}

	.button {
		display: inline-block;
		background: var(--color-theme-1);
		color: white;
		padding: 1rem 2rem;
		border-radius: 4px;
		text-decoration: none;
		font-weight: bold;
		transition: background 0.2s;
	}

	.button:hover {
		background: var(--color-theme-2);
	}
</style>
