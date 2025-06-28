#!/bin/bash
#
# Gerenciamento do Docker para LaTeX
#

# Configurações do Docker
readonly COMPOSE_FILE="${PROJECT_ROOT}/config/docker/docker-compose.yml"
# readonly DOCKERFILE_DIR="${PROJECT_ROOT}/config/docker/devcontainer"  # Não utilizada atualmente

# Verifica se o Docker está disponível
check_docker() {
    if ! command -v docker &> /dev/null; then
        log_error "Docker não encontrado. Por favor, instale o Docker."
        return 1
    fi

    if ! docker info &> /dev/null; then
        log_error "Docker não está executando. Por favor, inicie o Docker."
        return 1
    fi

    log_debug "Docker está disponível e executando"
    return 0
}

# Verifica se o Docker Compose está disponível
check_docker_compose() {
    if ! command -v docker-compose &> /dev/null && ! docker compose version &> /dev/null; then
        log_error "Docker Compose não encontrado. Por favor, instale o Docker Compose."
        return 1
    fi

    log_debug "Docker Compose está disponível"
    return 0
}

# Executa comando docker-compose
docker_compose() {
    local cmd="$1"
    shift

    if command -v docker-compose &> /dev/null; then
        docker-compose -f "$COMPOSE_FILE" "$cmd" "$@"
    else
        docker compose -f "$COMPOSE_FILE" "$cmd" "$@"
    fi
}

# Verifica se o container está executando
is_container_running() {
    local container_name="${1:-$(get_config container_name)}"

    if docker ps --filter "name=${container_name}" --filter "status=running" | grep -q "$container_name"; then
        return 0
    else
        return 1
    fi
}

# Verifica se o container existe (executando ou parado)
container_exists() {
    local container_name="${1:-$(get_config container_name)}"

    if docker ps -a --filter "name=${container_name}" | grep -q "$container_name"; then
        return 0
    else
        return 1
    fi
}

# Inicia o ambiente Docker
start_docker_env() {
    log_info "Iniciando ambiente Docker..."

    if ! check_docker || ! check_docker_compose; then
        return 1
    fi

    if is_container_running; then
        log_success "Container já está executando"
        return 0
    fi

    if docker_compose up -d; then
        log_success "Ambiente Docker iniciado com sucesso"

        # Aguarda o container ficar saudável
        wait_for_container_health
        return 0
    else
        log_error "Falha ao iniciar o ambiente Docker"
        return 1
    fi
}

# Para o ambiente Docker
stop_docker_env() {
    log_info "Parando ambiente Docker..."

    if docker_compose down; then
        log_success "Ambiente Docker parado com sucesso"
        return 0
    else
        log_error "Falha ao parar o ambiente Docker"
        return 1
    fi
}

# Reconstrói o ambiente Docker
rebuild_docker_env() {
    log_info "Reconstruindo ambiente Docker..."

    if docker_compose down && docker_compose build --no-cache && docker_compose up -d; then
        log_success "Ambiente Docker reconstruído com sucesso"
        wait_for_container_health
        return 0
    else
        log_error "Falha ao reconstruir o ambiente Docker"
        return 1
    fi
}

# Aguarda o container ficar saudável
wait_for_container_health() {
    local container_name="${1:-$(get_config container_name)}"
    local max_attempts=30
    local attempt=0

    log_info "Aguardando container ficar saudável..."

    while [[ $attempt -lt $max_attempts ]]; do
        if docker exec "$container_name" pdflatex --version &> /dev/null; then
            log_success "Container está saudável"
            return 0
        fi

        sleep 2
        ((attempt++))
        show_progress "Verificando saúde do container" "$attempt" "$max_attempts"
    done

    log_error "Container não ficou saudável após $max_attempts tentativas"
    return 1
}

# Executa comando no container
docker_exec() {
    local container_name="${1:-$(get_config container_name)}"
    shift

    if ! is_container_running "$container_name"; then
        log_error "Container $container_name não está executando"
        return 1
    fi

    docker exec -it "$container_name" "$@"
}

# Executa comando no container sem tty (útil para scripts)
docker_exec_no_tty() {
    local container_name="${1:-$(get_config container_name)}"
    shift

    if ! is_container_running "$container_name"; then
        log_error "Container $container_name não está executando"
        return 1
    fi

    docker exec "$container_name" "$@"
}

# Mostra logs do container
show_docker_logs() {
    local container_name="${1:-$(get_config container_name)}"
    local follow="${2:-false}"

    if [[ "$follow" == "true" ]]; then
        docker logs -f "$container_name"
    else
        docker logs "$container_name"
    fi
}

# Status do Docker
docker_status() {
    echo "=== Status do Docker ==="

    if check_docker; then
        echo "✓ Docker está disponível"
        docker version --format "Versão: {{.Server.Version}}"
    else
        echo "✗ Docker não está disponível"
        return 1
    fi

    if check_docker_compose; then
        echo "✓ Docker Compose está disponível"
    else
        echo "✗ Docker Compose não está disponível"
        return 1
    fi

    local container_name
    container_name="$(get_config container_name)"
    if is_container_running "$container_name"; then
        echo "✓ Container $container_name está executando"
    elif container_exists "$container_name"; then
        echo "⚠ Container $container_name existe mas não está executando"
    else
        echo "✗ Container $container_name não existe"
    fi

    return 0
}
