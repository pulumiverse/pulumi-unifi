import argparse
import json
import os
import ssl
import sys

from urllib import request, error, parse


base_url = ""


def _send_request(request_path, payload=None):
    incescure_ssl_ctx = ssl.create_default_context()
    incescure_ssl_ctx.check_hostname = False
    incescure_ssl_ctx.verify_mode = ssl.CERT_NONE

    data = None
    if payload is not None:
        data = json.dumps(payload).encode("utf-8")

    req = request.Request(parse.urljoin(base_url, request_path), data=data)

    if payload is not None:
        req.add_header("Content-Type", "application/json")

    return request.urlopen(req, context=incescure_ssl_ctx)


def is_controller_setup_required():
    path = "manage/account/login"
    response = _send_request(path)
    return not response.url.endswith(path)


def setup_local_admin(email, username, password):
    payload = {
        "cmd": "add-default-admin",
        "email": email,
        "name": username,
        "x_password": password
    }
    _send_request("api/cmd/sitemgr", payload)


def finalize_setup():
    payload = {"cmd": "set-installed"}
    _send_request("api/cmd/system", payload)


def parse_arguments():
    parser = argparse.ArgumentParser(description="Setup unifi controller.")
    parser.add_argument("--host", dest="host",
                    default=os.environ.get("UNIFI_HOST", "https://unifi:8443/"),
                    help="Host of the Unifi controller. Can be set using env 'UNIFI_HOST'")
    parser.add_argument("--email", dest="email",
                    default=os.environ.get("UNIFI_EMAIL", "pulumi-unifi@pulumi.com"),
                    help="E-Mail of admin user. Can be set using env 'UNIFI_EMAIL'")
    parser.add_argument("--username", dest="username",
                    default=os.environ.get("UNIFI_USERNAME", "pulumi-unifi-admin"),
                    help="Username of admin user. Can be set using env 'UNIFI_USERNAME'")
    parser.add_argument("--password", dest="password",
                    default=os.environ.get("UNIFI_PASSWORD", "pulumi-unifi-password"),
                    help="Password of admin user. Can be set using env 'UNIFI_PASSWORD'")
    return parser.parse_args()


if __name__ == "__main__":
    try:
        args = parse_arguments()
        base_url = args.host
        if is_controller_setup_required():
            print("Bootstrapping Unifi controller")
            setup_local_admin(args.email, args.username, args.password)
            finalize_setup()
        else:
            print("Unifi controller has already been setup")
    except error.URLError:
        print("Error setting up Unifi controller")
        sys.exit(1)
