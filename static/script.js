document.addEventListener('DOMContentLoaded', function () {
    // Function to fetch Lorem Ipsum message from the API
    function fetchLoremIpsum() {
        fetch('/api/lorem')
            .then(response => response.text())
            .then(data => {
                document.getElementById('message').textContent = data;
            })
            .catch(error => console.error('Error fetching lorem ipsum:', error));
    }

    // Load initial Lorem Ipsum message
    fetchLoremIpsum();

    // Event listener for Generate button
    document.getElementById('generateButton').addEventListener('click', function () {
        fetchLoremIpsum();
    });
});
