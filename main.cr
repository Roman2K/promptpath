require "yaml"

mapping = Mapping.new
mapping.load_env_val ENV.fetch("PROMPTPATH", "")
File.open ENV["HOME"] + "/.config/promptpath.yaml" do |f|
  mapping.load YAML.parse(f)
end

case ARGV.size
when 0
  mapping.each do |path, name|
    puts "#{name}\t#{path}"
  end
when 1
  puts mapping.shorten ARGV[0]
else
  raise ArgumentError.new "usage: #{File.basename PROGRAM_NAME} [path]"
end

class Mapping < Hash(String, String)
  def load(config)
    fill(config, nil)

    @re = %r{^(#{
      keys.
        sort_by { |k| [-k.count('/'), -k.size] }.
        map { |k| Regex.escape k }.
        join "|"
    })(?:/(.+)|$)}
  end

  def load_env_val(s)
    s.split(":").each do |kv|
      name, abs = kv.split("=", 2)
      self[abs] = name
    end
  end

  private def fill(config, root)
    config.as_h.each do |key, value|
      path = key.as_s?
      path = nil if path == "_"

      abs = [root, path].compact.join("/").gsub("~", ENV.fetch("HOME"))
      !abs.empty? || raise "empty path"

      if name = value.as_s?
        self[abs] = name
      else
        fill value, abs
      end
    end
  end

  def shorten(path)
    path =~ @re || return path
    short = self[$1]
    rel = $2? || return short
    "#{color "90"}#{short}/#{color "0"}#{rel}"
  end

  private def color(attrs)
    "\\[\e[#{attrs}m\\]"
  end
end
