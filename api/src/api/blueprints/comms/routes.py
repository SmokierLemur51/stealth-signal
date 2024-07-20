from quart import Blueprint

comms = Blueprint('public', __name__, template_folder="templates/public", url_prefix="/")

@comms.route("/<str:group>")