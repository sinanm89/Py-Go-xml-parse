from flask import Flask, Response, make_response
import xml.etree.ElementTree as ET

import requests
from werkzeug.exceptions import BadRequest

app = Flask(__name__)


@app.route('/<id>')
@app.route('/')
def vast_parser(id=65617):
    """
    Takes in the xml vast 2.0 request and relays it as vast 3.0

    :param id:
    :return:
    """
    url = "http://ad4.liverail.com/?LR_PUBLISHER_ID=%s&LR_SCHEMA" % str(id)
    print url
    vast_request = requests.get(url)
    if vast_request.status_code != 200:
        return BadRequest()

    vast = ET.fromstring(vast_request.content)

    if vast.attrib.get('content') == 'error':
        return BadRequest()
    vast.attrib['version'] = "3.0"

    # Generator will generate the singular item list
    creatives = list(vast.iter('Creatives'))[0]
    creatives.attrib["offset"]="00:00:08"
    return Response(ET.tostring(vast), mimetype='text/xml')
    # return vast

if __name__ == '__main__':
    app.debug = True
    app.run(host='0.0.0.0', port=5000, threaded=True)
