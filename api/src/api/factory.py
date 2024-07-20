from dotenv import load_dotenv
from quart import Quart

# settings.py env loading
load_dotenv()

def create_app(**config_overrides) -> Quart:
    app = Quart(__name__, static_url_path="/static")
    
    # .env file config
    app.config.from_pyfile("settings.py")

    # config obj
    from .config import Config
    app.config.from_object(Config)

    # config overrides
    app.config.update(config_overrides)
    
    # blueprints
    from .blueprints.comms.routes import comms
    app.register_blueprint(comms)

    # database
    from .extensions import db
    db.init_app(app)

    return app
