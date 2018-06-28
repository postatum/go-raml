# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
from .User2_0 import User2_0
from .unhandled_api_error import UnhandledAPIError
from .unmarshall_error import UnmarshallError


class UriService:
    def __init__(self, client):
        self.client = client

    def uri_byUsers_id_get(self, users_id, headers=None, query_params=None, content_type="application/json"):
        """
        It is method for GET /uri/{users-id}
        """
        if query_params is None:
            query_params = {}

        uri = self.client.base_url + "/uri/" + users_id
        resp = self.client.get(uri, None, headers, query_params, content_type)
        try:
            if resp.status_code == 200:
                return User2_0(resp.json()), resp

            message = 'unknown status code={}'.format(resp.status_code)
            raise UnhandledAPIError(response=resp, code=resp.status_code,
                                    message=message)
        except ValueError as msg:
            raise UnmarshallError(resp, msg)
        except UnhandledAPIError as uae:
            raise uae
        except Exception as e:
            raise UnmarshallError(resp, e.message)
