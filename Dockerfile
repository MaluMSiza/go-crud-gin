# Usa a imagem oficial do PostgreSQL
FROM postgres:15

# Define as variáveis padrão para inicialização do banco
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=go_crud_db

# Copia o script de inicialização SQL, se necessário
# (opcional, remova se não for necessário)
COPY ./init.sql /docker-entrypoint-initdb.d/

# Expondo a porta padrão do PostgreSQL
EXPOSE 5432
