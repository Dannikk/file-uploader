import streamlit as st
from io import StringIO
from main import get_client, paginator, printer
from streamlit_ace import st_ace


def on_click_upload_button():
    st.session_state.buttonOFF = True


upload_tab, download_tab = st.tabs(["Upload", "Download"])

client = None
if 'client' not in st.session_state:
    client = get_client()
    st.session_state['client'] = client
    print("Successfully created client")
else:
    client = st.session_state['client']
    print("Using existing client")

with upload_tab:
    uploaded_file = st.file_uploader("Выберите файл для загрузки")

    button_upload_disabled = True
    st.session_state.buttonOFF = True
    if uploaded_file is not None:
        if uploaded_file.name in paginator(client):
            st.error("Файл с таким именем уже существует")
        else:
            button_upload_disabled = False
            st.session_state.buttonOFF = False
            if st.button("Загрузить",
                         on_click=on_click_upload_button,
                         disabled=st.session_state.buttonOFF,
                         ):
                st.toast('Подождите, загружаю на облако...')
                result = client.upload_fileobj(uploaded_file, 'bucket-0', uploaded_file.name)
                button_upload_disabled = True
                # st.toast('Файл успешно загружен')
                print(uploaded_file, dir(uploaded_file))
                st.success('Файл успешно загружен')

            if uploaded_file:
                try:
                    content = st_ace(value=StringIO(uploaded_file.getvalue().decode("utf-8")).read(),
                                     min_lines=12,
                                     theme="solarized_dark",
                                     max_lines=25,
                                     height=300,
                                     readonly=True,
                                     )
                except BaseException as e:
                    print(e)

with download_tab:
    files = paginator(client)
    option = st.selectbox(
        "Выберите файл для скачивания",
        files,
        index=None,
        placeholder="Select file...",
    )
    file2download = None
    button_disabled = True
    # print("disabled before: ", button_disabled)
    # print("option: ", option)
    if option:
        file2download = option
        button_disabled = False
    # print("disabled: ", button_disabled)
    st.download_button('Скачать',
                       'some_file.txt',
                       disabled=button_disabled,
                       file_name=file2download,
                       on_click=printer,
                       )
    st.write(files)
