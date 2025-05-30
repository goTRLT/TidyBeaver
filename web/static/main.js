document.addEventListener("DOMContentLoaded", function() {
    const logTableBody = document.getElementById("logTableBody");
    const filterInput = document.getElementById("filterInput");
    const sortSelect = document.getElementById("sortSelect");

    let logs = [];

    async function fetchLogs() {
        const response = await fetch('/logs');
        logs = await response.json();
        displayLogs(logs);
    }

    function displayLogs(logs) {
        logTableBody.innerHTML = "";
        logs.forEach(log => {
            const row = document.createElement("tr");
            row.innerHTML = `
                <td>${new Date(log.timestamp).toLocaleString()}</td>
                <td>${log.service}</td>
                <td>${log.level}</td>
                <td>${log.message}</td>
            `;
            logTableBody.appendChild(row);
        });
    }

    function filterLogs() {
        const filterValue = filterInput.value.toLowerCase();
        const filteredLogs = logs.filter(log => 
            log.service.toLowerCase().includes(filterValue) ||
            log.level.toLowerCase().includes(filterValue) ||
            log.message.toLowerCase().includes(filterValue)
        );
        displayLogs(filteredLogs);
    }

    function sortLogs() {
        const sortBy = sortSelect.value;
        const sortedLogs = [...logs].sort((a, b) => {
            if (sortBy === "timestamp") {
                return new Date(a.timestamp) - new Date(b.timestamp);
            } else if (sortBy === "service") {
                return a.service.localeCompare(b.service);
            } else if (sortBy === "level") {
                return a.level.localeCompare(b.level);
            }
            return 0;
        });
        displayLogs(sortedLogs);
    }

    filterInput.addEventListener("input", filterLogs);
    sortSelect.addEventListener("change", sortLogs);

    fetchLogs();
});