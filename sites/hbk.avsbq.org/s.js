// Function to fetch the site index from an external file
async function fetchSiteIndex() {
    const response = await fetch('https://hbk.avsbq.org/s.json'); // Replace with the correct path to your JSON file
    const data = await response.json();
    return data;
}

// Function to filter search results
function filterResults(query, siteIndex) {
    return siteIndex.filter(item => item.toLowerCase().includes(query.toLowerCase()));
}

// Function to display search results
function displayResults(results) {
    const searchResults = document.getElementById('searchResults');
    searchResults.innerHTML = '';

    // Slice the results to show a maximum of 5 items
    results.slice(0, 4).forEach(result => {
        const resultItem = document.createElement('div');
        resultItem.className = 'resultItem';
        resultItem.innerHTML = `<a href="/${result}">${result}</a>`;
        searchResults.appendChild(resultItem);
    });
}

// Event listener for search input
document.getElementById('searchBar').addEventListener('input', async function() {
    const query = this.value;
    const searchResults = document.getElementById('searchResults');
    
    // Clear search results if input is empty
    if (query === '') {
        searchResults.innerHTML = '';
        return;
    }

    const siteIndex = await fetchSiteIndex();
    const results = filterResults(query, siteIndex);
    displayResults(results);
});
