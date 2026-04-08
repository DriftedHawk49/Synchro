# ---------- Runtime stage ----------
FROM nginx:1.29.6

RUN curl -L https://github.com/yt-dlp/yt-dlp/releases/download/2026.03.13/yt-dlp_linux \
    -o /usr/local/bin/yt-dlp \
    && chmod +x /usr/local/bin/yt-dlp
# Create a non-root user (optional but civilized)



WORKDIR /app

COPY ./YoutubeSynchroniser .



ENTRYPOINT ["./YoutubeSynchroniser"]
