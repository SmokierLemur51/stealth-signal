import os
basedir = os.path.abspath(os.path.dirname(__file__))


class Config(object):
    DEBUG = False
    TESTING = False
    CSRF_ENABLED = True
    SECRET_KEY = 'this-really-needs-to-be-changed'
    # SQLALCHEMY_DATABASE_URI="sqlite:///{}.db".format(os.environ['SQLITE_DB'])
    QUART_DB_DATABASE_URL = "postgresql://{}:{}@localhost:5432/{}".format(
        os.environ["PG_USER"], os.environ["PG_PASS"], os.environ["PG_DATABASE"])

class ProductionConfig(Config):
    DEBUG = False


class StagingConfig(Config):
    DEVELOPMENT = True
    DEBUG = True


class DevelopmentConfig(Config):
    DEVELOPMENT = True
    DEBUG = True


class TestingConfig(Config):
    TESTING = True
