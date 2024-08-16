async function fetchPaths() {
    const response = await fetch('s.json');
    const paths = await response.json();
    return paths;
}

async function goToRandomPage() {
    const paths = await fetchPaths();
    const randomPath = paths[Math.floor(Math.random() * paths.length)];
    window.location.href = randomPath;
}

