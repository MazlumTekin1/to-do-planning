document.getElementById('getBtn').addEventListener('click', function() {
    fetch('http://localhost:3000/tasks')
    .then(response => response.json())
    .then(data => {
        displayMaxWeeks(data.max_weeks);
        createTable(data.data);
    })
    .catch(error => {
        console.error('Hata:', error);
    });
});

function createTable(data) {
    const tableContainer = document.getElementById('tableContainer');
    tableContainer.innerHTML = ''; // Önceki tabloyu temizle

    const table = document.createElement('table');
    const headerRow = table.insertRow();
    const headers = ['Developer', 'Görev', 'Süre', 'Zorluk'];
    headers.forEach(headerText => {
        const th = document.createElement('th');
        th.textContent = headerText;
        headerRow.appendChild(th);
    });

    for (const developer in data) {
        data[developer].forEach(task => {
            const row = table.insertRow();
            row.insertCell().textContent = developer;
            row.insertCell().textContent = task.Name;
            row.insertCell().textContent = task.Duration;
            row.insertCell().textContent = task.Difficulty;
        });
    }

    tableContainer.appendChild(table);
}

function displayMaxWeeks(maxWeeks) {
    const maxWeeksDiv = document.getElementById('maxWeeks');
    maxWeeksDiv.textContent = `Developer'lar bu işi minimum ${maxWeeks} Hafta içerisinde bitirebilir.`;
}



document.getElementById('postBtn').addEventListener('click', function() {
    fetch('http://localhost:3000/tasks/add', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
    })
    .then(response => response.json())
    .then(data => {
        showNotification(data.message);
    })
    .catch(error => {
        console.error('Hata:', error);
    });
});

function showNotification(message) {
    if (!('Notification' in window)) {
        console.error('Tarayıcınız bildirimleri desteklemiyor.');
        return;
    }

    Notification.requestPermission().then(function(permission) {
        if (permission === 'granted') {
            new Notification(message);
        } else {
            console.error('Kullanıcı bildirim iznini vermedi.');
        }
    });
}



