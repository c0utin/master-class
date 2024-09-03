class Block:
  def __init__(self, data, hash, last_hash):
      self.data = data
      self.hash = hash
      self.last_hash = last_hash

def lightning_hash(data):
  return data + "*superSeguro*"

class Blockchain:
  def __init__(self):
      genesis = Block("gen-data", "gen-hash", "gen-lastHash")
      self.chain = [genesis]

  def add_block(self, data):
      last_block = self.chain[-1]
      last_hash = last_block.hash
      hash = lightning_hash(data + last_hash)
      block = Block(data, hash, last_hash)
      self.chain.append(block)

blockchain_poggers = Blockchain()
blockchain_poggers.add_block("one")
blockchain_poggers.add_block("two")
blockchain_poggers.add_block("three")

for block in blockchain_poggers.chain:
  print(f"Data: {block.data}, Hash: {block.hash}, LastHash: {block.last_hash}")
  print("---------------------------------------")
