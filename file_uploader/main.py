import os
import boto3
import yaml


def get_client():
    cfg = load_config()
    session = boto3.session.Session()
    s3 = session.client(
        service_name='s3',
        endpoint_url='https://storage.yandexcloud.net',
        aws_access_key_id=cfg["aws_access_key_id"],
        aws_secret_access_key=cfg["aws_secret_access_key"],
    )

    # ooo = s3.get_paginator('list_objects').paginate(Bucket='bucket-0')
    # print(ooo)
    # for item in ooo:
    #     print(item)
    # for key in s3.list_objects(Bucket='bucket-0')['Contents']:
    #     print(key, key['Key'])

    return s3


def upload(client, path: str):
    print(client.upload_file(path, 'bucket-0', path.split('/')[-1]))


def download(client, file):
    print(client.download_file('bucket-0', file, 'new_downloaded.txt'))


def paginator(client):
    files = []
    for key in client.list_objects(Bucket='bucket-0')['Contents']:
        file_name = key['Key']
        # print(file_name)
        files.append(file_name)
    return files


def load_config():
    with open("./config/config.yaml", 'r') as file:
        config = yaml.safe_load(file)
    return config


def printer(*args):
    print("I'm in printer", args)


if __name__ == '__main__':
    print(os.listdir("../"))
    client = get_client()


    # upload(client, "./test_file_other.txt")
    paginator(client)
    # download(client)
