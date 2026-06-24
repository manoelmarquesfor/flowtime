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
        const response = await rawFetch(url, {
            headers: {
                'Content-Type': 'application/json',
                ...options.headers
            },
            ...options
        });
        const data = await response.json().catch(() => ({}));

        if (response.ok) {
            return data;
        }

        if (response.status === 401) {
            showToast("Sessão expirada. Redirecionando para login...", 'error');
            window.location.href = "/login";
            return;
        }

        throw new Error(data.detail || `Erro HTTP ${response.status}`);

    } catch (error) {
        showToast(error.message, 'error');
        throw error;
    }
}


async function rawFetch(url, options = {}) {
    return fetch(url, {
        credentials: "include",
        ...options
    });
}

async function checkAuth() {
    try {
        const res = await rawFetch("/auth/me");

        if (!res.ok) {
            window.location.href = "/login";
            return;
        }

        document.body.style.display = "block";

    } catch {
        window.location.href = "/login";
    }
}


