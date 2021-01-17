import etcd3
import time

if __name__ == '__main__':
    client = etcd3.client(host="10.0.6.239", port=12379)
    v, _ = client.get("some_key")
    print(v)  # b'some_value_20210116'

    with client.lock("resource/1") as lock:
        print(lock.key)
        time.sleep(10)
