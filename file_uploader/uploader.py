import streamlit as st
import pandas as pd
from io import StringIO
from main import get_client, paginator
from boto3 import s3

upload_tab, download_tab = st.tabs(["Upload", "Download"])

client = None
if 'client' not in st.session_state:
    client = get_client()
    print(client)
    st.session_state['client'] = client
    print("Successfully created client")
else:
    client = st.session_state['client']
    print("Using existing client")


with upload_tab:
    uploaded_file = st.file_uploader("Choose a file")
    if uploaded_file is not None:
        client.upload_fileobj(uploaded_file, 'bucket-0', uploaded_file.name)
        print(uploaded_file, dir(uploaded_file))
        st.write(uploaded_file)
        st.write(uploaded_file.name)

with download_tab:
    files = paginator(client)
    st.write(files)
    st.write("Download file")
    st.write_stream(files)
    st.table(files)
