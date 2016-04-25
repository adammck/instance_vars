# InstanceVars

This thing fetches [metadata][] from GCE instances, and dumps it in a format
suitable for setting environment variables. Don't use this, it's probably an
awful idea. I'm just giving it a try.

## Usage

    instance_vars > /etc/environment

## License

MIT.

[metadata]: https://cloud.google.com/compute/docs/metadata#custom
