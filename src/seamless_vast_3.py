from flask import Flask, Response, make_response, request
import xml.etree.ElementTree as ET # phone home

import requests
from werkzeug.exceptions import BadRequest

app = Flask(__name__)


@app.route('/')
@app.route('/<id>')
def vast_parser(id=None):
    """
    Takes in the xml vast 2.0 request and relays it as vast 3.0
    either :arg publisher_id: or :param id: exists
    :return:
    """
    if request.args.get('publisher_id'):
        publisher_id = request.args.get('publisher_id')
    elif id:
        publisher_id = id
    else:
        return BadRequest()

    url = "http://ad4.liverail.com/?LR_PUBLISHER_ID=%s&LR_SCHEMA" % str(publisher_id)
    print url
    vast_request = requests.get(url)
    if vast_request.status_code != 200:
        return BadRequest()

    vast = ET.fromstring(vast_request.content)

    if vast.attrib.get('content') == 'error':
        return BadRequest()
    vast.attrib['version'] = "3.0"

    # Generator will generate the singular item list
    creatives = list(vast.iter('Linear'))[0]
    creatives.attrib["skip_offset"]="00:00:08"
    return Response(ET.tostring(vast), mimetype='text/xml')
    # return vast

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, threaded=True)
