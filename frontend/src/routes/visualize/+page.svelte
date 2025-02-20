<script lang="ts">
    import { onMount } from 'svelte';

    interface QueryResult {
        columns: string[];
        rows: any[];
        executionTime: string;
        rowCount: number;
    }

    let activeTab = 'table';
    let selectedVisualization = 'bar';
    
    // Mock data for demonstration
    let queryResult: QueryResult = {
        columns: ['id', 'name', 'value', 'date'],
        rows: [
            { id: 1, name: 'Item 1', value: 100, date: '2024-03-01' },
            { id: 2, name: 'Item 2', value: 150, date: '2024-03-02' },
            { id: 3, name: 'Item 3', value: 75, date: '2024-03-03' },
        ],
        executionTime: '0.123s',
        rowCount: 3
    };

    const visualizationTypes = [
        { value: 'bar', label: 'Bar Chart' },
        { value: 'line', label: 'Line Chart' },
        { value: 'pie', label: 'Pie Chart' },
        { value: 'scatter', label: 'Scatter Plot' }
    ];

    function downloadCSV() {
        const headers = queryResult.columns.join(',');
        const rows = queryResult.rows.map(row => 
            queryResult.columns.map(col => row[col]).join(',')
        ).join('\n');
        
        const csv = `${headers}\n${rows}`;
        const blob = new Blob([csv], { type: 'text/csv' });
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = 'query_results.csv';
        a.click();
        window.URL.revokeObjectURL(url);
    }

    function renderVisualization() {
        // TODO: Implement visualization rendering using a charting library
        console.log('Rendering visualization:', selectedVisualization);
    }
</script>

<svelte:head>
    <title>QB | Visualize</title>
    <meta name="description" content="Visualize query results" />
</svelte:head>

<div class="container">
    <div class="header">
        <h1>Query Results</h1>
        <div class="stats">
            <span>Execution Time: {queryResult.executionTime}</span>
            <span>Rows: {queryResult.rowCount}</span>
        </div>
    </div>

    <div class="tabs">
        <button 
            class:active={activeTab === 'table'} 
            on:click={() => activeTab = 'table'}
        >
            Table View
        </button>
        <button 
            class:active={activeTab === 'visual'} 
            on:click={() => activeTab = 'visual'}
        >
            Visualization
        </button>
    </div>

    {#if activeTab === 'table'}
        <div class="table-container">
            <div class="table-actions">
                <button on:click={downloadCSV}>Download CSV</button>
            </div>
            <table>
                <thead>
                    <tr>
                        {#each queryResult.columns as column}
                            <th>{column}</th>
                        {/each}
                    </tr>
                </thead>
                <tbody>
                    {#each queryResult.rows as row}
                        <tr>
                            {#each queryResult.columns as column}
                                <td>{row[column]}</td>
                            {/each}
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
    {:else}
        <div class="visualization-container">
            <div class="visualization-controls">
                <select bind:value={selectedVisualization}>
                    {#each visualizationTypes as vizType}
                        <option value={vizType.value}>{vizType.label}</option>
                    {/each}
                </select>
                <button on:click={renderVisualization}>Update Chart</button>
            </div>
            <div class="chart-area">
                <!-- Chart will be rendered here -->
                <p>Chart placeholder for {selectedVisualization}</p>
            </div>
        </div>
    {/if}
</div>

<style>
    .container {
        max-width: 1200px;
        margin: 0 auto;
        padding: 2rem;
    }

    .header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 2rem;
    }

    .stats {
        display: flex;
        gap: 1rem;
        font-size: 0.9rem;
        color: var(--color-text-2);
    }

    .tabs {
        display: flex;
        gap: 1rem;
        margin-bottom: 2rem;
    }

    .tabs button {
        padding: 0.5rem 1rem;
        border: none;
        background: none;
        cursor: pointer;
        border-bottom: 2px solid transparent;
    }

    .tabs button.active {
        border-bottom: 2px solid var(--color-theme-1);
        color: var(--color-theme-1);
    }

    .table-container {
        overflow-x: auto;
    }

    .table-actions {
        margin-bottom: 1rem;
    }

    table {
        width: 100%;
        border-collapse: collapse;
    }

    th, td {
        padding: 0.75rem;
        text-align: left;
        border-bottom: 1px solid var(--color-bg-2);
    }

    th {
        background: var(--color-bg-1);
        font-weight: bold;
    }

    .visualization-container {
        padding: 2rem;
        background: var(--color-bg-1);
        border-radius: 8px;
    }

    .visualization-controls {
        display: flex;
        gap: 1rem;
        margin-bottom: 2rem;
    }

    .chart-area {
        min-height: 400px;
        display: flex;
        align-items: center;
        justify-content: center;
        background: var(--color-bg-2);
        border-radius: 4px;
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

    select {
        padding: 0.5rem;
        border: 1px solid var(--color-bg-2);
        border-radius: 4px;
    }
</style> 