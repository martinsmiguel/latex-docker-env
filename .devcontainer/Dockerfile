# Use a imagem oficial mínima do TeX Live
FROM texlive/texlive:latest

# Otimiza instalação e tamanho da imagem
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        # Dependências essenciais
        perl \
        git \
        # Ferramentas de compilação
        latexmk \
        biber \
        # Suporte a códigos (minted)
        python3-pygments \
        # Visualizador leve (opcional para debug)
        xzdec \
        evince \
        # Limpeza pós-instalação
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /tmp/*

# Configura usuário não-root para segurança
ARG USERNAME=latexuser
ARG USER_UID=1001
ARG USER_GID=1001

# Cria o usuário com UID diferente para evitar conflito
RUN groupadd --gid $USER_GID $USERNAME && \
    useradd --uid $USER_UID --gid $USER_GID -m $USERNAME

# Instala pacotes LaTeX adicionais via tlmgr (com tentativas de mirror)
RUN tlmgr option repository https://mirror.ctan.org/systems/texlive/tlnet && \
    tlmgr update --self --verify-repo=none || \
    (tlmgr option repository https://ctan.org/tex-archive/systems/texlive/tlnet && \
     tlmgr update --self --verify-repo=none) && \
    tlmgr install --verify-repo=none \
        enumitem \
        fancyhdr \
        hyperref \
        xcolor \
    && tlmgr path add

# Configura ambiente de trabalho
USER $USERNAME
WORKDIR /workspace
VOLUME ["/workspace"]

# Health check para monitorar o serviço
HEALTHCHECK --interval=30s --timeout=5s \
    CMD pdflatex --version || exit 1