<script lang="ts">
	import { onMount } from 'svelte';
	import cytoscape from 'cytoscape';

	let fileInput: HTMLInputElement;
	let cy;
	let searchTerm = '';
	let confidenceFilter = 0;
	let nodesRaw = [];
	let edgesRaw = [];
	let selectedNode = null;
	let relatedEdges = [];

	async function handleUpload() {
		const formData = new FormData();
		formData.append("file", fileInput.files[0]);

		const res = await fetch('http://localhost:8080/api/upload', {
			method: 'POST',
			body: formData
		});

		if (res.ok) {
			alert("✅ Dataset cargado correctamente.");
			location.reload();
		} else {
			alert("❌ Error al subir el archivo.");
		}
	}

	onMount(async () => {
		const resEntities = await fetch('http://localhost:8080/api/entities');
		const resRelations = await fetch('http://localhost:8080/api/relations');
		nodesRaw = await resEntities.json();
		edgesRaw = await resRelations.json();
		renderGraph();
	});

	function renderGraph() {
		if (!cy) {
			cy = cytoscape({
				container: document.getElementById('cy'),
				style: [
					{
						selector: 'node',
						style: {
							'background-color': '#2563eb',
							'label': 'data(id)',
							'font-size': '12px',
							'color': '#ffffff',
							'font-family': 'IBM Plex Mono',
							'text-outline-color': '#1e3a8a',
							'text-outline-width': 1
						}
					},
					{
						selector: 'edge',
						style: {
							'width': 2,
							'line-color': '#94a3b8',
							'target-arrow-color': '#64748b',
							'target-arrow-shape': 'triangle',
							'curve-style': 'bezier',
							'label': 'data(label)',
							'font-size': '10px',
							'color': '#cbd5e1'
						}
					}
				],
				layout: { name: 'cose', animate: true }
			});
		} else {
			cy.elements().remove();
		}

		const filteredEdges = edgesRaw.filter(e => e.Confidence >= confidenceFilter);
		const usedNodes = new Set(filteredEdges.flatMap(e => [e.Source, e.Target]));

		const filteredNodes = nodesRaw.filter(n => {
			const id = n.Nombre || n.IP || n.Hash;
			return usedNodes.has(id) && id.toLowerCase().includes(searchTerm.toLowerCase());
		});

		const cyNodes = filteredNodes.map(n => ({
			data: { id: n.Nombre || n.IP || n.Hash }
		}));

		const cyEdges = filteredEdges.map(r => ({
			data: {
				source: r.Source,
				target: r.Target,
				label: r.Type,
				confidence: r.Confidence,
				source_info: r.SourceInfo
			}
		}));

		cy.add([...cyNodes, ...cyEdges]);
		cy.layout({ name: 'cose', animate: true }).run();

		cy.on('tap', 'node', function (evt) {
			selectedNode = evt.target.id();
			relatedEdges = cy.edges().filter((e) => e.data().source === selectedNode).map(e => e.data());
		});
	}
</script>

<style>
#cy {
	height: 80vh;
	width: 100%;
	background: rgba(255, 255, 255, 0.05);
	backdrop-filter: blur(12px);
	border-radius: 1rem;
	box-shadow: 0 0 20px rgba(0, 0, 0, 0.25);
	margin-top: 2rem;
}
</style>

<div class="min-h-screen bg-white/10 text-slate-100 px-8 py-6 font-mono">
	<h1 class="text-4xl font-bold text-blue-400 mb-4">🛰 Glasya-Labolas</h1>
	<p class="mb-6 text-slate-300">Análisis de relaciones tácticas</p>

	<form
		on:submit|preventDefault={handleUpload}
		class="bg-white/10 backdrop-blur px-4 py-3 rounded border border-slate-500 mb-4"
	>
		<h2 class="text-slate-200 text-sm mb-1">Subir nuevo CSV</h2>
		<input type="file" bind:this={fileInput} accept=".csv" class="mb-2 text-white" />
		<button
			type="submit"
			class="bg-blue-600 px-3 py-1 text-white text-sm rounded hover:bg-blue-500"
		>
			Cargar dataset
		</button>
	</form>

	<div class="flex gap-4 mb-4">
		<input
			bind:value={searchTerm}
			on:input={renderGraph}
			placeholder="Buscar entidad (IP, nombre, hash)"
			class="bg-white/10 backdrop-blur px-4 py-2 rounded border border-slate-500 text-white w-full"
		/>
		<input
			type="range"
			min="0"
			max="100"
			step="10"
			bind:value={confidenceFilter}
			on:input={renderGraph}
			class="w-1/3"
		/>
		<span class="text-sm text-slate-300">Confianza ≥ {confidenceFilter}</span>
	</div>

	<div class="flex gap-6">
		<div class="w-3/4">
			<div id="cy"></div>
		</div>

		<div class="w-1/4 bg-white/10 backdrop-blur rounded p-4 border border-slate-500 text-sm">
			{#if selectedNode}
				<h2 class="text-xl text-blue-300 font-bold mb-2">🧠 Nodo seleccionado</h2>
				<p><strong>ID:</strong> {selectedNode}</p>
				<p><strong>Relaciones salientes:</strong> {relatedEdges.length}</p>
				<ul class="mt-2 space-y-2">
					{#each relatedEdges as edge}
						<li class="border-l-2 pl-2 border-blue-400">
							<strong>{edge.label}</strong><br />
							➡ {edge.target}<br />
							Confianza: {edge.confidence}
							<small class="block text-slate-400 italic">Fuente: {edge.source_info}</small>
						</li>
					{/each}
				</ul>
			{:else}
				<p class="text-slate-400">Haz clic en un nodo para ver detalles.</p>
			{/if}
		</div>
	</div>
</div>