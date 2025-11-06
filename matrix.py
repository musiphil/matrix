#!/usr/bin/python3
import argparse
import fileinput
import sys


class Matrix(object):

  def __init__(self):
    self.rows = []
    self.cols = []
    self.body = {}

  def add(self, r, c, b):
    if r not in self.rows:
      self.rows.append(r)
    if c not in self.cols:
      self.cols.append(c)
    self.body[r, c] = '' if b is None else b

  def print(self, file, sep, missing):
    rows, cols = self.rows, self.cols
    print('', *cols, sep=sep, file=file)

    for r in rows:
      cells = [self.body.get((r, c), missing) for c in cols]
      print(r, *cells, sep=sep, file=file)


if __name__ == '__main__':
  parser = argparse.ArgumentParser()
  parser.add_argument('files', nargs='*')
  parser.add_argument('--input-sep')
  parser.add_argument('--output-sep', default='\t')
  parser.add_argument('--missing', default='')
  args = parser.parse_args()

  m = Matrix()
  for line in fileinput.input(args.files):
    line = line.rstrip()
    items = line.split(args.input_sep, maxsplit=2)
    if len(items) <= 1:
      r, = items
      c = ''
      b = None
    elif len(items) <= 2:
      r, c = items
      b = None
    else:
      r, c, b = items
    m.add(r, c, b)
  m.print(sys.stdout, sep=args.output_sep, missing=args.missing)
