# Base Ubuntu server
FROM ubuntu:24.04

# Cài tool cần thiết
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        wget \
        gnupg \
        ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Thêm GPG key và repo Grafana
RUN mkdir -p /etc/apt/keyrings/ && \
    wget -q -O - https://apt.grafana.com/gpg.key | gpg --dearmor | tee /etc/apt/keyrings/grafana.gpg > /dev/null && \
    echo "deb [signed-by=/etc/apt/keyrings/grafana.gpg] https://apt.grafana.com stable main" | tee /etc/apt/sources.list.d/grafana.list

# Update và cài alloy
RUN apt-get update && \
    apt-get install -y alloy && \
    rm -rf /var/lib/apt/lists/*

# Tạo thư mục storage và fix quyền cho user alloy (UID 473)
RUN mkdir -p /var/lib/alloy/data && \
    chown -R 473:473 /var/lib/alloy/data && \
    chmod 755 /var/lib/alloy/data

# COPY config.alloy từ context build vào container và fix quyền
COPY config.alloy /etc/alloy/config.alloy
RUN chown 473:473 /etc/alloy/config.alloy && \
    chmod 644 /etc/alloy/config.alloy

# Expose UI debug
EXPOSE 12345

# Command mặc định
CMD ["alloy", "run", "--server.http.listen-addr=0.0.0.0:12345", \
     "--storage.path=/var/lib/alloy/data", "/etc/alloy/config.alloy"]
