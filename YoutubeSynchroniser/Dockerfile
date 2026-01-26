# ---------- Runtime stage ----------
FROM alpine:3.19

# Create a non-root user (optional but civilized)
RUN adduser -D -g '' appuser

WORKDIR /app

COPY ./YoutubeSynchroniser .

USER appuser

ENTRYPOINT ["./YoutubeSynchroniser"]
