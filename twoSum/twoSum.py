#!/usr/bin/env python

class problems(object):
    def twoSum(self,nums,target):
        lookUP = {}
        for i,num in enumerate(nums):
            if target - num in lookUP:
                return [lookUP[target - num], i]
            lookUP[num] = i
