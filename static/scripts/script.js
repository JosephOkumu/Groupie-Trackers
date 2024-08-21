
document.addEventListener('DOMContentLoaded', () => {
    // Handle search form submission
    const searchForm = document.querySelector('form[action="/search"]');
    if (searchForm) {
        searchForm.addEventListener('submit', function(event) {
            event.preventDefault();
            const query = this.querySelector('input[name="query"]').value.trim();
            if (query) {
                fetch(`/search?query=${encodeURIComponent(query)}`)
                    .then(response => response.json())
                    .then(data => {
                        if (data && data.length > 0) {
                            displaySearchResults(data);
                        } else {
                            alert('No artists found.');
                        }
                    })
                    .catch(error => {
                        console.error('Error fetching search results:', error);
                        alert('Failed to fetch search results.');
                    });
            } else {
                alert('Please enter a search query.');
            }
        });
    }
});

function displaySearchResults(artists) {
    const resultsContainer = document.querySelector('.artists-grid');
    if (resultsContainer) {
        resultsContainer.innerHTML = ''; // Clear previous results
        artists.forEach(artist => {
            const artistCard = document.createElement('div');
            artistCard.className = 'artist-card';
            artistCard.innerHTML = `
                <a href="/artist?id=${artist.ID}">
                    <img src="${artist.Image}" alt="Image of ${artist.Name}" onerror="this.onerror=null;this.src='/static/images/placeholder.jpg';">
                    <h2>${artist.Name}</h2>
                </a>
                <a href="/artist?id=${artist.ID}" class="view-details-link">View Details</a>
            `;
            resultsContainer.appendChild(artistCard);
        });
    }
}