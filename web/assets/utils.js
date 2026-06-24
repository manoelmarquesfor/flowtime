// Container de Toasts injetado dinamicamente
const toastContainer = document.createElement('div');
toastContainer.id = 'toast-container';
document.body.appendChild(toastContainer);

function showToast(message, type = 'info') {
    const toast = document.createElement('div');
    toast.className = `toast ${type}`;
    toast.innerHTML = `
        <span>${message}</span>
        <span style="cursor:pointer; margin-left:15px; opacity:0.7" onclick="this.parentElement.remove()">✕</span>
    `;
    toastContainer.appendChild(toast);
    setTimeout(() => {
        toast.style.animation = 'fadeOutToast 0.3s ease forwards';
        setTimeout(() => toast.remove(), 300);
    }, 2000);
}

// Wrapper padronizado para Fetch (reduz repetição nos forms)
async function apiFetch(url, options = {}) {
    try {
        const response = await fetch(url, options);
        const data = await response.json().catch(() => ({}));
        if (!response.ok) {
            throw new Error(data.detail || `Erro HTTP ${response.status}`);
        }
        return data;
    } catch (error) {
        showToast(error.message, 'error');
        throw error;
    }
}