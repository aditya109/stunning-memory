FROM python
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY main.py ./
EXPOSE 3000
ENTRYPOINT [ "python" ]
CMD [ "main.py" ]
