<script lang="ts">
	import { sandboxMode } from '$lib/stores/sandbox';

	interface Connection {
		id: string;
		name: string;
		type: string;
		host: string;
		port: string;
		database: string;
		username: string;
	}

	let currentConnection: Connection | null = null;
	let selectedTable = '';
	let columns: string[] = [];
	let conditions: Array<{column: string, operator: string, value: string}> = [];
	let joins: Array<{table: string, type: string, condition: string}> = [];
	let loading = true;

	// Mock data for demonstration
	const tables = ['users', 'orders', 'products'];
	const availableColumns = ['id', 'name', 'email', 'created_at'];
	const operators = ['=', '!=', '>', '<', '>=', '<=', 'LIKE', 'IN'];

	async function loadConnection() {
		loading = true;
		try {
			const response = await fetch('http://localhost:3000/api/connections', {
				headers: {
					'X-Sandbox-Mode': $sandboxMode ? 'true' : 'false'
				}
			});
			const connections = await response.json();
			currentConnection = connections[0] || null;
		} catch (error) {
			console.error('Failed to load connection:', error);
			currentConnection = null;
		} finally {
			loading = false;
		}
	}

	// Subscribe to sandbox mode changes
	$: {
		if ($sandboxMode !== undefined) {
			loadConnection();
		}
	}

	function addCondition() {
		conditions = [...conditions, { column: '', operator: '=', value: '' }];
	}

	function addJoin() {
		joins = [...joins, { table: '', type: 'INNER', condition: '' }];
	}

	function removeCondition(index: number) {
		conditions = conditions.filter((_, i) => i !== index);
	}

	function removeJoin(index: number) {
		joins = joins.filter((_, i) => i !== index);
	}

	function generateQuery() {
		// TODO: Implement query generation logic
	}
</script>

<svelte:head>
	<title>QB | Query Builder</title>
	<meta name="description" content="Build SQL queries visually" />
</svelte:head>

<div class="container">
	<h1>Query Builder</h1>

	<div class="header-controls">
		<div class="toggle-container">
			<label class="toggle">
				<input
					type="checkbox"
					bind:checked={$sandboxMode}
				/>
				<span class="slider">
					<span class="left label">Real</span>
					<span class="right label">Sandbox</span>
				</span>
			</label>
		</div>
	</div>

	<div class="section connection-info">
		{#if loading}
			<div class="loading">Loading connection...</div>
		{:else if currentConnection}
			<div class="current-connection">
				<h2>Current Connection</h2>
				<p>{currentConnection.name} ({currentConnection.type})</p>
				<p class="connection-details">{currentConnection.host}:{currentConnection.port}/{currentConnection.database}</p>
			</div>
		{:else}
			<div class="no-connection">
				<h2>No Active Connection</h2>
				{#if $sandboxMode}
					<p>Using sandbox mode with mock connections.</p>
				{:else}
					<p>No real connections configured yet.</p>
				{/if}
				<div class="connection-options">
					<a href="/connections" class="button primary">
						<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
							<path d="M6 8v-2c0-2.2 1.8-4 4-4h4c2.2 0 4 1.8 4 4v2"/>
							<line x1="12" y1="12" x2="12" y2="16"/>
							<circle cx="12" cy="20" r="2"/>
						</svg>
						{$sandboxMode ? 'Configure Real Connection' : 'Configure a Connection'}
					</a>
				</div>
			</div>
		{/if}
	</div>

	{#if currentConnection}
		<div class="section">
			<h2>Table Selection</h2>
			<select bind:value={selectedTable}>
				<option value="">Select Table</option>
				{#each tables as table}
					<option value={table}>{table}</option>
				{/each}
			</select>
		</div>

		<div class="section">
			<h2>Columns</h2>
			<div class="columns-grid">
				{#each availableColumns as column}
					<label>
						<input
							type="checkbox"
							bind:group={columns}
							value={column}
						/>
						{column}
					</label>
				{/each}
			</div>
		</div>

		<div class="section">
			<h2>Conditions</h2>
			<button on:click={addCondition}>Add Condition</button>
			{#each conditions as condition, i}
				<div class="condition-row">
					<select bind:value={condition.column}>
						<option value="">Select Column</option>
						{#each availableColumns as column}
							<option value={column}>{column}</option>
						{/each}
					</select>
					<select bind:value={condition.operator}>
						{#each operators as op}
							<option value={op}>{op}</option>
						{/each}
					</select>
					<input type="text" bind:value={condition.value} placeholder="Value" />
					<button on:click={() => removeCondition(i)}>Remove</button>
				</div>
			{/each}
		</div>

		<div class="section">
			<h2>Joins</h2>
			<button on:click={addJoin}>Add Join</button>
			{#each joins as join, i}
				<div class="join-row">
					<select bind:value={join.type}>
						<option value="INNER">INNER JOIN</option>
						<option value="LEFT">LEFT JOIN</option>
						<option value="RIGHT">RIGHT JOIN</option>
						<option value="FULL">FULL JOIN</option>
					</select>
					<select bind:value={join.table}>
						<option value="">Select Table</option>
						{#each tables as table}
							<option value={table}>{table}</option>
						{/each}
					</select>
					<input type="text" bind:value={join.condition} placeholder="ON condition" />
					<button on:click={() => removeJoin(i)}>Remove</button>
				</div>
			{/each}
		</div>

		<div class="section">
			<h2>Generated Query</h2>
			<div class="query-preview">
				<pre>SELECT * FROM table WHERE 1=1;</pre>
			</div>
			<button on:click={generateQuery}>Generate Query</button>
		</div>
	{/if}
</div>

<style>
	.container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 2rem;
	}

	.section {
		margin-bottom: 2rem;
		padding: 1.5rem;
		border: 1px solid var(--color-border);
		border-radius: 8px;
		background: var(--color-bg-1);
	}

	.loading {
		text-align: center;
		padding: 2rem;
		color: var(--color-text-2);
	}

	.no-connection {
		text-align: center;
		padding: 2rem;
	}

	.no-connection h2 {
		margin-bottom: 1rem;
		color: var(--color-text);
	}

	.no-connection p {
		margin-bottom: 2rem;
		color: var(--color-text-2);
	}

	.connection-options {
		display: flex;
		gap: 1rem;
		justify-content: center;
		align-items: center;
	}

	.button {
		display: inline-flex;
		align-items: center;
		gap: 0.5rem;
		padding: 0.75rem 1.5rem;
		border-radius: 6px;
		font-weight: 500;
		text-decoration: none;
		transition: all 0.2s;
	}

	.button.primary {
		background: var(--color-theme-1);
		color: white;
	}

	.button.primary:hover {
		background: var(--color-theme-2);
	}

	.current-connection {
		padding: 1rem;
		border-radius: 6px;
	}

	.current-connection h2 {
		margin-bottom: 0.5rem;
		color: var(--color-text);
	}

	.connection-details {
		font-family: monospace;
		color: var(--color-text-2);
	}

	svg {
		stroke-width: 2;
	}

	.columns-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
		gap: 1rem;
	}

	.condition-row, .join-row {
		display: grid;
		grid-template-columns: 1fr 1fr 2fr auto;
		gap: 1rem;
		margin-bottom: 1rem;
		align-items: center;
	}

	.query-preview {
		background: #f5f5f5;
		padding: 1rem;
		border-radius: 4px;
		margin: 1rem 0;
	}

	button {
		background: var(--color-theme-1);
		color: white;
		border: none;
		padding: 0.5rem 1rem;
		border-radius: 4px;
		cursor: pointer;
	}

	button:hover {
		background: var(--color-theme-2);
	}

	select, input {
		padding: 0.5rem;
		border: 1px solid #ccc;
		border-radius: 4px;
	}

	.header-controls {
		display: flex;
		justify-content: flex-end;
		margin-bottom: 1rem;
	}

	.toggle-container {
		display: flex;
		align-items: center;
		gap: 0.5rem;
	}

	.toggle {
		position: relative;
		display: inline-block;
		width: 160px;
		height: 40px;
		cursor: pointer;
	}

	.toggle input {
		display: none;
	}

	.slider {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background-color: var(--color-bg-2);
		border: 2px solid var(--color-theme-1);
		border-radius: 20px;
		transition: 0.4s;
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 0 10px;
	}

	.slider:before {
		position: absolute;
		content: "";
		height: 32px;
		width: 76px;
		left: 2px;
		bottom: 2px;
		background-color: var(--color-theme-1);
		border-radius: 16px;
		transition: 0.4s;
	}

	.label {
		color: var(--color-theme-1);
		font-weight: 500;
		z-index: 1;
		transition: 0.4s;
	}

	.label.right {
		color: var(--color-bg-2);
	}

	input:checked + .slider:before {
		transform: translateX(76px);
	}

	input:checked + .slider .label.left {
		color: var(--color-theme-1);
	}

	input:checked + .slider .label.right {
		color: var(--color-bg-2);
	}
</style> 