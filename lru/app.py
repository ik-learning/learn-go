#! /usr/bin/env python

"""step 0"""
class LRU():

    """step 0"""
    def __init__(self, capacity):
        self.cache = {}
        self.key_order = []
        self.capacity = capacity

    def __setitem__(self, key, value):
        self.cache[key] = value
        """step 4"""
        self._process(key)

    def add(self, key, value):
        self.cache[key] = value

    """step 2"""
    def get(self, key):
        "try and catch step 3.i"
        try:
            value = self.cache[key]
        except KeyError:
            raise KeyError, key
        self._process(key)
        return value

    def _process(self, key):
        if key in self.key_order:
            self.key_order.remove(key)
        print(self.key_order)
        self.key_order.insert(0, key)
        print(self.key_order)
        if len(self.key_order) > self.capacity:
            remove = self.key_order[self.capacity]
            del self.cache[remove]
            self.key_order.remove(remove)



"""step 1"""
if __name__ == '__main__':
    import unittest

    class TestLRU(unittest.TestCase):
        def test_should_only_have_most_used_items(self):
            cache = LRU(capacity=4)
            cache['v1'] = 1
            cache['v2'] = 23
            cache['v8'] = 28
            cache['v9'] = 34
            cache['v10'] = 89

            "try and catch step 3"
            self.assertRaises(KeyError, cache.get, "v1")
            print("v2")
            self.assertEqual(cache.get('v2'), 23)
            print("v3")
            self.assertEqual(cache.get('v10'), 89)



    unittest.main()