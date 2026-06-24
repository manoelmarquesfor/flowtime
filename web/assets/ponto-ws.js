const IMGS_PATH = "/imgs/";
const DEFAULT_IMAGE = "/assets/default.svg";

function normalizeImageUrl(rawImage) {
    if (!rawImage) return DEFAULT_IMAGE;
    const value = rawImage.trim();
    const isAbsolute = value.startsWith("http") || value.startsWith("data:") || value.startsWith("blob:");
    if (isAbsolute || value.startsWith(IMGS_PATH)) return value;
    if (/^\/?[\w.-]+\.(jpg|jpeg|png|gif|webp)$/i.test(value)) {
        return `${IMGS_PATH}${value.replace(/^\//, '')}`;
    }
    return value;
}

class PontoWS {
    constructor(callbacks, options = {}) {
        this.callbacks = callbacks; // { onStatus, onEntidade, onClear }
        this.silent = options.silent || false; // usado para silenciar o toast como no ponto_matricula onde o feedback já é dado via API
        this.connect();
    }

    connect() {
        const protocol = window.location.protocol === "https:" ? "wss" : "ws";
        this.ws = new WebSocket(`${protocol}://${window.location.host}/ws/funcionario`);

        this.ws.onopen = () => {
            this.callbacks.onStatus("Aguardando batida...", true);
        };

        this.ws.onclose = () => {
            this.callbacks.onStatus("Conexão perdida. Reconectando...", false);
            showToast("Conexão com o leitor perdida", "warning");
            setTimeout(() => this.connect(), 2000);
        };

        this.ws.onerror = () => {
            this.callbacks.onStatus("Erro na conexão WebSocket", false);
        };

        this.ws.onmessage = (event) => {
            try {
                const msg = JSON.parse(event.data);

                // Tratamento explícito pelo "type" acordado para o Backend
                // Fallback para verificação de chaves ("duck typing") se o backend ainda não mandou
                if (msg.type === "clear") {
                    this.callbacks.onClear();
                }
                else if (msg.type === "detail" || ('detail' in msg)) {
                    if (!this.silent) {
                        // Erros reportados ativamente pelo backend via WebSocket
                        showToast(msg.detail || "Erro desconhecido informado pelo servidor.", "error");
                    }
                }
                else if (msg.type === "entidade" || ('imagem' in msg)) {
                    // Sucesso / Leitura Efetuada
                    if (msg.imagem && msg.imagem.trim() !== "") {
                        msg.imagem = normalizeImageUrl(msg.imagem);
                        this.callbacks.onEntidade(msg);
                    }
                }
                else {
                    showToast("Formato de mensagem WebSocket não reconhecido.", "warning");
                }
            } catch (err) {
                showToast("Erro ao processar dados do servidor.", "error");
            }
        };
    }
}