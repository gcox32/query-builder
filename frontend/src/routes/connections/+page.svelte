<script lang="ts">
    import { onMount } from 'svelte';

    interface Connection {
        id: string;
        name: string;
        type: string;
        host: string;
        port: string;
        database: string;
        username: string;
        password: string;
    }

    let connections: Connection[] = [];
    let showNewConnectionForm = false;
    let testingConnection = false;
    let testResult: { success: boolean; message: string; time: string } | null = null;
    let loading = true;
    let newConnection: Connection = {
        id: crypto.randomUUID(),
        name: '',
        type: 'postgresql',
        host: '',
        port: '',
        database: '',
        username: '',
        password: ''
    };

    const databaseTypes = [
        { value: 'postgresql', label: 'PostgreSQL' },
        { value: 'mysql', label: 'MySQL' },
        { value: 'mongodb', label: 'MongoDB' },
        { value: 'sqlite', label: 'SQLite' }
    ];

    onMount(() => {
        loadConnections();
    });

    async function loadConnections() {
        loading = true;
        try {
            const response = await fetch('http://localhost:3000/api/connections');
            connections = await response.json();
        } catch (error) {
            console.error('Failed to load connections:', error);
        } finally {
            loading = false;
        }
    }

    async function addConnection() {
        connections = [...connections, { ...newConnection }];
        
        // TODO: Implement actual connection storage
        showNewConnectionForm = false;
        newConnection = {
            id: crypto.randomUUID(),
            name: '',
            type: 'postgresql',
            host: '',
            port: '',
            database: '',
            username: '',
            password: ''
        };
    }

    function deleteConnection(id: string) {
        connections = connections.filter(conn => conn.id !== id);
    }

    async function testConnection(connection: Connection) {
        testingConnection = true;
        testResult = null;
        
        try {
            const response = await fetch('http://localhost:3000/api/connections/test', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(connection),
            });
            
            testResult = await response.json();
        } catch (error) {
            testResult = {
                success: false,
                message: 'Failed to test connection: Network error',
                time: '0s'
            };
        } finally {
            testingConnection = false;
        }
    }
</script>

<svelte:head>
    <title>QB | Connections</title>
    <meta name="description" content="Manage database connections" />
</svelte:head>

<div class="container">
    <div class="header">
        <h1>Database Connections</h1>
        <button on:click={() => showNewConnectionForm = true}>New Connection</button>
    </div>

    {#if loading}
        <div class="loading">Loading connections...</div>
    {:else if connections.length === 0 && !showNewConnectionForm}
        <div class="no-connections">
            <p>No connections configured yet.</p>
            <button on:click={() => showNewConnectionForm = true}>Add Your First Connection</button>
        </div>
    {/if}

    {#if showNewConnectionForm}
        <div class="connection-form">
            <h2>New Connection</h2>
            <form on:submit|preventDefault={addConnection}>
                <div class="form-group">
                    <label for="name">Connection Name</label>
                    <input type="text" id="name" bind:value={newConnection.name} required />
                </div>

                <div class="form-group">
                    <label for="type">Database Type</label>
                    <select id="type" bind:value={newConnection.type}>
                        {#each databaseTypes as dbType}
                            <option value={dbType.value}>{dbType.label}</option>
                        {/each}
                    </select>
                </div>

                <div class="form-group">
                    <label for="host">Host</label>
                    <input type="text" id="host" bind:value={newConnection.host} required />
                </div>

                <div class="form-group">
                    <label for="port">Port</label>
                    <input type="text" id="port" bind:value={newConnection.port} required />
                </div>

                <div class="form-group">
                    <label for="database">Database Name</label>
                    <input type="text" id="database" bind:value={newConnection.database} required />
                </div>

                <div class="form-group">
                    <label for="username">Username</label>
                    <input type="text" id="username" bind:value={newConnection.username} required />
                </div>

                <div class="form-group">
                    <label for="password">Password</label>
                    <input type="password" id="password" bind:value={newConnection.password} required />
                </div>

                <div class="button-group">
                    <button type="submit">Save Connection</button>
                    <button type="button" class="secondary" on:click={() => showNewConnectionForm = false}>
                        Cancel
                    </button>
                </div>
            </form>
        </div>
    {/if}

    <div class="connections-list">
        {#each connections as connection}
            <div class="connection-card">
                <div class="connection-info">
                    <h3>{connection.name}</h3>
                    <p>{connection.type} • {connection.host}:{connection.port}</p>
                </div>
                <div class="connection-actions">
                    <button 
                        class="secondary" 
                        on:click={() => testConnection(connection)}
                        disabled={testingConnection}
                    >
                        {#if testingConnection}
                            Testing...
                        {:else}
                            Test Connection
                        {/if}
                    </button>
                    <button class="danger" on:click={() => deleteConnection(connection.id)}>
                        Delete
                    </button>
                </div>
                {#if testResult}
                    <div class="test-result" class:success={testResult.success} class:error={!testResult.success}>
                        <p>{testResult.message}</p>
                        <small>Time: {testResult.time}</small>
                    </div>
                {/if}
            </div>
        {/each}
    </div>
</div>

<style>
    .container {
        max-width: 800px;
        margin: 0 auto;
        padding: 2rem;
    }

    .header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 2rem;
    }

    .connection-form {
        background: var(--color-bg-1);
        padding: 2rem;
        border-radius: 8px;
        margin-bottom: 2rem;
    }

    .form-group {
        margin-bottom: 1rem;
    }

    .form-group label {
        display: block;
        margin-bottom: 0.5rem;
        font-weight: bold;
    }

    .form-group input,
    .form-group select {
        width: 100%;
        padding: 0.5rem;
        border: 1px solid #ccc;
        border-radius: 4px;
    }

    .button-group {
        display: flex;
        gap: 1rem;
        margin-top: 1rem;
    }

    .connection-card {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 1rem;
        background: var(--color-bg-1);
        border-radius: 4px;
        margin-bottom: 1rem;
    }

    .connection-actions {
        display: flex;
        gap: 0.5rem;
    }

    button {
        padding: 0.5rem 1rem;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        background: var(--color-theme-1);
        color: white;
    }

    button.secondary {
        background: var(--color-bg-2);
        color: var(--color-text);
    }

    button.danger {
        background: #dc3545;
    }

    button:hover {
        opacity: 0.9;
    }

    .test-result {
        margin-top: 1rem;
        padding: 0.5rem;
        border-radius: 4px;
        font-size: 0.9rem;
    }

    .test-result.success {
        background: #d4edda;
        color: #155724;
        border: 1px solid #c3e6cb;
    }

    .test-result.error {
        background: #f8d7da;
        color: #721c24;
        border: 1px solid #f5c6cb;
    }

    .test-result small {
        display: block;
        margin-top: 0.25rem;
        opacity: 0.8;
    }

    button:disabled {
        opacity: 0.7;
        cursor: not-allowed;
    }

    .loading {
        text-align: center;
        padding: 2rem;
        color: var(--color-text-2);
    }

    .no-connections {
        text-align: center;
        padding: 3rem;
        background: var(--color-bg-1);
        border-radius: 8px;
        margin-top: 2rem;
    }

    .no-connections p {
        margin-bottom: 1rem;
        color: var(--color-text-2);
    }

    .header-actions {
        display: flex;
        align-items: center;
        gap: 1rem;
    }
</style> 