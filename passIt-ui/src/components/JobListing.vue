<script setup lang="ts">
import jobData from '@/jobs.json';
import { ref, defineProps } from 'vue';
import { type Job } from '@/models/Jobs'
import JobListings from '@/components/JobListings.vue';

defineProps({
    limit: Number,
    showAll: {
        type: Boolean,
        default: false
    },
});

const jobs = ref<Job[]>(jobData.jobs);

</script>

<template>
    <section class="bg-blue-50 px-4 py-10">
        <div class="container-xl lg:container m-auto">
            <h2 class="text-3xl font-bold text-green-500 mb-6 text-center">
                Browse Jobs
            </h2>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                <JobListings v-for="job in jobs.slice(0, limit || jobs.length)" :key="job.id" :job="job" />
            </div>
        </div>
    </section>

    <section v-if="showAll" class="m-auto max-w-lg my-10 px-6">
      <a
        href="/jobs"
        class="block bg-black text-white text-center py-4 px-6 rounded-xl hover:bg-gray-700"
        >View All Jobs</a
      >
    </section>
    
</template>