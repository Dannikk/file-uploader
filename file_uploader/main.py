import os
import boto3
import yaml


def get_client():
    session = boto3.session.Session()
    s3 = session.client(
        service_name='s3',
        endpoint_url='https://storage.yandexcloud.net',
        aws_access_key_id=os.getenv("aws_access_key_id"),
        aws_secret_access_key=os.getenv("aws_secret_access_key"),
    )

    return s3


def upload(client, path: str):
    print(client.upload_file(path, 'bucket-0', path.split('/')[-1]))


def paginator(client):
    for key in client.list_objects(Bucket='bucket-0')['Contents']:
        print(key['Key'])


def load_config():
    with open("../config/config_example.yaml", 'r') as file:
        config = yaml.safe_load(file)
    return config


if __name__ == '__main__':
    print(os.listdir("../"))
    client = get_client()

    upload(client, "../test_file.txt")
    # paginator(client)
