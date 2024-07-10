function generateArt() {
    const text = document.getElementById('text').value;
    const banner = document.querySelector('input[name="banner"]:checked').value;
    fetch('/ascii-art-live', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `text=${encodeURIComponent(text)}&banner=${encodeURIComponent(banner)}`,
    })
    .then(response => response.text())
    .then(data => {
        document.getElementById('ascii-art-result').innerHTML = data.replace(/\n/g, '\r\n');
    })
    .catch(error => {
        console.error('Error:', error);
    });
}