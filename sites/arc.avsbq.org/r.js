async function fetchPaths() {
    const response = await fetch('s.json');
    const paths = await response.json();
    return paths;
}

async function goToRandomPage() {
    const paths = await fetchPaths();

    // Filter out paths that include "misc/personal/school"
    const filteredPaths = paths.filter(path => !path.includes('misc/personal/school'));

    if (filteredPaths.length === 0) {
        console.error("No valid paths available.");
        return;
    }

    const randomPath = filteredPaths[Math.floor(Math.random() * filteredPaths.length)];
    window.location.href = randomPath;
}

