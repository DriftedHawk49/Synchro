FROM ubuntu:noble
RUN apt update && apt install ffmpeg wget -y
RUN wget https://github.com/yt-dlp/yt-dlp/releases/download/2025.12.08/yt-dlp_linux && mv yt-dlp_linux yt-dlp && chmod +x yt-dlp && mv yt-dlp /usr/local/bin/yt-dlp
RUN adduser -D -g '' appuser

WORKDIR /app
RUN mkdir /app/songs && chown adduser /app/songs

COPY ./YoutubeDownloader .

USER appuser

ENTRYPOINT ["./YoutubeDownloader"]
