"""
This module provides solutions to leetcode questions.
"""


def merge(nums1: list[int], m: int, nums2: list[int], n: int) -> list[int]:
    """Merges two sorted list. time nlongn"""
    nums1 = nums1[:m] + nums2
    nums1.sort()
    return nums1


def remove_element(nums: list[int], val: int) -> int:
    """remove element from given list"""
    counter = 0
    for value in nums:
        if value != val:
            nums[counter] = value
            counter += 1

    return counter


def main():
    """Main function where all the functions will go through"""
    print("This is the main function.")
    print("___________________________________________________________________")
    print("Merge Sorted Array")
    merge_list1 = [1, 2, 3, 0, 0, 0]
    merge_list2 = [2, 5, 6]
    merge_list1 = merge(merge_list1, 3, merge_list2, 3)
    print(merge_list1)
    print("___________________________________________________________________")
    print("Removed Element")
    remove_list = [3, 1, 2, 4, 3, 5, 8]
    merge_list1 = remove_element(remove_list, 3)
    print(merge_list1)
    print("___________________________________________________________________")


if __name__ == "__main__":
    main()
