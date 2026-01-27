# FROM ubuntu:noble
# RUN apt update && apt install ffmpeg wget -y
# RUN wget https://github.com/yt-dlp/yt-dlp/releases/download/2025.12.08/yt-dlp_linux && mv yt-dlp_linux yt-dlp && chmod +x yt-dlp && mv yt-dlp /usr/local/bin/yt-dlp

# WORKDIR /app
# RUN mkdir /app/songs

# COPY ./YoutubeDownloader .

# ENTRYPOINT ["./YoutubeDownloader"]


# -------------------------
# Stage 2: yt-dlp
# -------------------------


# -------------------------
# Stage 3: final runtime
# -------------------------
FROM linuxserver/ffmpeg:8.0.1

RUN curl -L https://github.com/yt-dlp/yt-dlp/releases/download/2025.12.08/yt-dlp_linux \
    -o /usr/local/bin/yt-dlp \
    && chmod +x /usr/local/bin/yt-dlp

# your Go binary
WORKDIR /app
RUN mkdir /app/songs

COPY ./YoutubeDownloader .

ENTRYPOINT ["./YoutubeDownloader"]
