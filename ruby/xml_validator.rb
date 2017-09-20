# frozen_string_literal: true

# <continent name="Europe">
#   <country name="Netherlands">
#   </country>
#   <country name="France">
#   </country>
# </continent>

# But what about?

# <continent name="Europe"> # leaf? false, next = country name="Netherlands"
# <country name="Netherlands"> # leaf true, next = continent
# </continent>
# </country>

# <continent name="Europe"> # leaf? false, next = country name="Netherlands"
# <country name="Netherlands"> # leaf true, next = continent
# </country>
# </country>
=begin
def well_formed?(tag[])
    int count = 0
    for(i=0;i<tag.length;i++){
        tag = tag[i]
        if(tag.isOpen) {
            count ++
        }
        else {
            count --
        }
        if count < 0 {
            return false
        }
    }

    return count == 0
    }
}

<country></coutnry>i
def well_formed2(tags)
    if(tags.first.isOpen) {
    node = new Node(name: tags.first)
    }else { return false}
    well_formed_recurse( (tags - tags.first), node)
end

def well_formed_recurse(tags, node)
    if( tags.isEmpty) {
        return node.isNil
    }
    if(tags.first.isClosed && node.name = tag.closed){
        node.remove
    }
    node.child = tags.first
    return well_formed_recurse( (tags - tags.first), node)
}

# Better solution would be to use a stack
=end

def well_formed(string)
  tags = parse_tags(string)
  stack = []
  tags.each do |t|
    if t.open?
      stack.push(t)
    else
      return false unless stack.pop&.closed_by?(t)
    end
  end
  stack.empty?
end

def parse_tags(string)
  string.scan(/<([^>]+)>/).map do |e|
    open = e.first[0] != "/"
    name = open ? e.first.split(" ").first : e.first[1..-1].split(" ").first
    Tag.new(name, open)
  end
end

class Tag
  attr_accessor :name, :open

  def initialize(name, open = true)
    self.name = name
    self.open = open
  end

  def open?
    open
  end

  def closed_by?(tag)
    tag.name == name && open? && !tag.open?
  end

  def ==(other)
    name == other.name && open == other.open
  end

  def to_s
    "<#{open? ? '' : '/'}#{name}>"
  end
end
