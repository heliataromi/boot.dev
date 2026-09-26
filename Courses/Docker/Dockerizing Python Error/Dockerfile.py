FROM debian:stable-slim
COPY main.py main.py
COPY books/ books/
CMD ["python", "main.py"]
