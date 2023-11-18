defmodule Day3 do
  def get_coords(index, steps, map_size) do
    {step_x, step_y} = steps

    {map_x, map_y} = map_size

    {rem(index * step_x, map_x), rem(index * step_y, map_y)}
  end

  def count_trees(input, map_size, steps) do
    {_, step_y} = steps

    input
    |> Enum.with_index()
    |> Enum.filter(fn {_, index} -> rem(index, step_y) end)
    |> Enum.map(fn {_, index} ->
      {x, y} = get_coords(index, steps, map_size)
      if(input |> Enum.at(y) |> String.at(x) == "#", do: 1, else: 0)
    end)
    |> Enum.sum()
  end
end

input = Path.absname("day3.txt", "./input/2020")
|> File.read!()
|> String.split("\n")

map_x = input |> List.first() |> String.length()
map_y = input |> length()

map_size = {map_x, map_y}

input
|> Day3.count_trees(map_size, {3, 1})
|> IO.puts()
